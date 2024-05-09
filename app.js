require("dotenv").config();

const mongoose = require("mongoose");
const express = require("express");
const fileUpload = require("express-fileupload");
const bodyParser = require("body-parser");

const newsRouter = require("./src/routes/news.routes");
const eventRouter = require("./src/routes/event.routes");
const userRouter = require("./src/routes/user.routes");

const app = express();

const urlencodedParser = bodyParser.urlencoded({
  extended: false,
});

app.use(express.static("public"));
app.use(express.json());
app.use(fileUpload());
app.use("/news", newsRouter);
app.use("/event", eventRouter);
app.use("/user", userRouter);

app.set("view engine", "hbs");
app.set("views", "public/views");

app.get("/edit-account", function (req, res) {
  res.render("edit-account.hbs");
});

app.get("/register", function (req, res) {
  res.render("register.hbs");
});

app.post("/register", urlencodedParser, async function (req, res) {
  try {
    if (!req.body) throw new Error("Request body is empty!");
    await fetch("http://localhost:3000/user", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...req.body, role: "user" }),
    })
      .then((response) => {
        if (!response.ok) throw new Error("Create user request was not ok");
        res.status(200).redirect("/login");
        // ИСПРАВИТЬ ЭТОТ КОСТЫЛЬ!
        return;
      })
      .catch((error) => {
        console.error(error);
      });
  } catch (err) {
    console.log(err);
  }
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
