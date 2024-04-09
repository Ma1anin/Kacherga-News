const express = require('express');
const path = require('path');

const port = 3000;
const app = express();
const router = express.Router();

app.use(express.static('public'));

app.use('/', function(_, res) {
  res.redirect('pages/main.html');
});

app.listen(port, () => console.log('Сервер запущен... (port: 5000)'));
