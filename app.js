var express = require('express');

const app = express();
const router = express.Router();

app.use(express.static('public'));

app.use('/', function (_, response) {
  response.redirect('/pages/main.html');
});

app.listen(5000, () => console.log('Сервер запущен... (port: 5000)'));
