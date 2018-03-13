// eslint-disable-next-line import/no-extraneous-dependencies
const config = require('config');
const session = require('express-session');
const passport = require('passport');
const { Strategy: FacebookStrategy } = require('passport-facebook');
const GoogleStrategy = require('passport-google-oauth20').Strategy;

passport.serializeUser((user, done) => {
  done(null, user);
});
passport.deserializeUser((user, done) => {
  done(null, user);
});

module.exports = server => {
  server.use(
    session({
      resave: true,
      saveUninitialized: true,
      secret: config.get('session.secret'),
      maxAge: 86400,
    })
  );
  server.use(passport.initialize());
  server.use(passport.session());

  function extractProfile(profile) {
    let imageUrl = '';
    if (profile.photos && profile.photos.length) {
      imageUrl = profile.photos[0].value;
    }
    return {
      id: profile.id,
      displayName: profile.displayName,
      image: imageUrl,
    };
  }

  passport.use(
    new GoogleStrategy(
      {
        clientID: config.get('google.clientID'),
        clientSecret: config.get('google.clientSecret'),
        callbackURL: config.get('google.callbackURL'),
      },
      (accessToken, refreshToken, profile, cb) => {
        cb(null, extractProfile(profile));
      }
    )
  );

  server.get(
    '/auth/google',
    (req, res, next) => {
      if (req.query.return) {
        req.session.oauth2return = req.query.return;
      }
      next();
    },
    passport.authenticate('google', { scope: ['email', 'profile'] })
  );

  server.get(
    '/auth/google/callback',
    passport.authenticate('google'),
    (req, res) => {
      const redirect = req.session.oath2return || '/';
      delete req.session.oauth2return;
      res.redirect(redirect);
    }
  );

  function authRequired(req, res, next) {
    if (!req.user) {
      req.session.oauth2return = req.originalUrl;
      return res.redirect('/login');
    }
    next();
  }

  // function addTemplateVariables(req, res, next) {
  //   res.locals.profile = req.user;
  //   res.locals.login = `/auth/login?return=${encodeURIComponent(
  //     req.originalUrl
  //   )}`;
  //   res.locals.logout = `/auth/logout?return=${encodeURIComponent(
  //     req.originalUrl
  //   )}`;
  //   next();
  // }

  passport.use(
    new FacebookStrategy(
      {
        clientID: config.get('facebook.clientID'),
        clientSecret: config.get('facebook.clientSecret'),
        callbackURL: config.get('facebook.callbackURL'),
        profileFields: ['id', 'age_range', 'email', 'name', 'gender'],
      },
      (accessToken, refreshToken, profile, cb) => {
        cb(null, profile);
        // User.findOrCreate({ facebookId: profile.id }, function(err, user) {
        //   return cb(err, user);
        // });
      }
    )
  );

  server.get(
    '/auth/facebook',
    passport.authenticate('facebook', {
      scope: ['user:email'],
    })
  );

  server.get(
    '/auth/facebook/callback',
    passport.authenticate('facebook', {
      failureRedirect: '/login',
    }),
    (req, res) => {
      res.send('OK');
      // res.json(user);
      // res.redirect('/');
    }
  );

  server.get('/protected', authRequired, (req, res) => {
    res.json(req.user);
  });

  return server;
};
