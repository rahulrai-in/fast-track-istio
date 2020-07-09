using System.Collections.Generic;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace legacy.Controllers
{
    [ApiController]
    [Route("books")]
    public class BooksController : ControllerBase
    {
        [HttpGet]
        public IEnumerable<Book> Get()
        {
            return JsonConvert.DeserializeObject<IList<Book>>(System.IO.File.ReadAllText("data.json"));
        }
    }
}