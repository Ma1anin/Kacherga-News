import request from "supertest";
import { app } from "./app";
  
it("should return correct user", function(done){
    request(app)
        .get("/user/1")
        .expect({
            login: "cortezz",
            password: "null",
            name: "Rodrigo Milk",
            role: "super-admin",
            picture: "null.jpg"
        })
        .end(done);
});
