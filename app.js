require("dotenv").config();

const mongoose = require("mongoose");
const express = require("express");
const fileUpload = require("express-fileupload");
const newsRouter = require("./src/routes/news.routes");
const eventRouter = require("./src/routes/event.routes");
const userRouter = require("./src/routes/user.routes");

const app = express();

app.use(express.static("public"));
app.use(express.json());
app.use(fileUpload());
app.use("/news", newsRouter);
app.use("/event", eventRouter);
app.use("/user", userRouter);

app.set("view engine", "hbs");
app.set("views", "public/views");

app.get("/editAccount", function (req, res) {
  res.render("editAccount.hbs");
});

app.get("/register", function (req, res) {
  res.render("register.hbs");
});

app.get("/login", function (req, res) {
  res.render("login.hbs");
});

app.get("/", function (req, res) {
  res.render("main.hbs");
});

async function startApp() {
  try {
    await mongoose.connect(process.env.DB_URI);
    app.listen(process.env.PORT, () =>
      console.log(`Server started on port ${process.env.PORT}`)
    );
  } catch (err) {
    console.log(err);
  }
}

startApp();
