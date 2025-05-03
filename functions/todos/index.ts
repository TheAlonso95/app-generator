import { Application, Router, Context } from "https://deno.land/x/oak/mod.ts";

const router = new Router();



// Delete a todo
router.delete("/todos/{id}", async (ctx: Context) => {
    

    
    ctx.response.status = 200;
    ctx.response.body = {};
    
});


// Get a todo by ID
router.get("/todos/{id}", async (ctx: Context) => {
    

    
    ctx.response.status = 200;
    ctx.response.body = {
  "id": 0,  "title": "example_string",  "completed": false
};
    
});


// Update a todo
router.put("/todos/{id}", async (ctx: Context) => {
    
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
