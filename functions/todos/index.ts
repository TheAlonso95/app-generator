import { Application, Router } from "https://deno.land/x/oak/mod.ts";

const router = new Router();



// Get all todos
router.get("/todos", async (ctx) => {
    

    
    ctx.response.status = 200;
    ctx.response.body = {};
    
});


// Create a new todo
router.post("/todos", async (ctx) => {
    
    const body = await ctx.request.body().value;
    console.log("Received:", body);
    

    
    ctx.response.status = 200;
    ctx.response.body = {
  "id": 0,  "title": "example_string",  "completed": false
};
    
});



const app = new Application();
app.use(router.routes());
app.use(router.allowedMethods());

console.log("Server running on http://localhost:8000");
await app.listen({ port: 8000 });
