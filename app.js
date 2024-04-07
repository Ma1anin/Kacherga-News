var express = require("express");

const app = express();

app.use(express.static("public"));

app.listen(7545, () => {
  console.log("Сервер запущен... (port: 7545)");
});
