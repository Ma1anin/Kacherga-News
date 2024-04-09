require('dotenv').config();

const express = require('express');
const path = require('path');

const app = express();
const router = express.Router();

app.use(express.static('public'));

app.use('/', function(_, res) {
  res.redirect('pages/main.html');
});

app.listen(process.env.PORT, () => console.log(`Server started on port ${process.env.PORT}`));
