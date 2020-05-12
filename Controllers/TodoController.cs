using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

using System.Text.Json;
using System.Text.Json.Serialization;

namespace TodoApi.Controllers
{
    [ApiController]
    [Route("[controller]")]
    [Produces("application/json")]
    public class TodoController : ControllerBase
    {
        private static readonly string[] UnfinishedTasks = new[]
        {
            "Feed the dogs", "Walk the dogs", "Finish demo"
        };

        private readonly ILogger<TodoController> _logger;

        public TodoController(ILogger<TodoController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public string Get()
        {
            TodoItem[] todos = new TodoItem[3];
            for (int i = 0; i < 3; ++i)
            {
                todos[i] = new TodoItem
                {
                    Complete = false,
                    Name = UnfinishedTasks[i],
                    Id = i
                };
            }

            return JsonSerializer.Serialize(todos);
        }
    }
}
