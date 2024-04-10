require('dotenv').config();

const express = require('express');
const path = require('path');
const newsRouter = require('./src/routes/news.routes');
const eventRouter = require('./src/routes/event.routes');
const userRouter = require('./src/routes/user.routes');

const app = express();

app.use(express.static('public'));

app.set('view engine', 'hbs');
app.set('views', 'public/views');

app.use(function (req, res) {
  res.render('main.hbs');
});

app.listen(process.env.PORT, () => console.log(`Server started on port ${process.env.PORT}`));
