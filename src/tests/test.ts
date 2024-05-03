import request from "supertest";
import { app } from "./app";
  
it("should return correct user", function(done){
    request(app)
        .get("/user/1")
        .expect(function(response){
            assert.deepStrictEqual(response.body, {
                login: "cortezz",
                name: "Rodrigo Milk",
                role: "super-admin",
                picture: "null.jpg"
            });
        })
        .end(done);
});
