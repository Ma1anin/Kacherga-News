require('dotenv').config();

const express = require('express');
const path = require('path');
const newsRouter = require('./src/routes/news.routes')

const app = express();

app.use(express.static('public'));

app.use('/', function(_, res) {
  res.redirect('pages/main.html');
});

app.listen(process.env.PORT, () => console.log(`Server started on port ${process.env.PORT}`));
