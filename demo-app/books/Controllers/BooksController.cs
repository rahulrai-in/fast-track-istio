using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace books.Controllers
{
    [ApiController]
    [Route("books")]
    public class BooksController : Controller
    {
        [HttpGet]
        public IEnumerable<Book> Get()
        {
            var books = JsonConvert.DeserializeObject<IList<Book>>(System.IO.File.ReadAllText("data.json"));
            foreach (var book in books)
            {
                var imageArray = System.IO.File.ReadAllBytes(book.Thumbnail);
                var base64ImageRepresentation = Convert.ToBase64String(imageArray);
                book.Thumbnail = base64ImageRepresentation;
                book.Summary = string.Empty;
            }

            return books;
        }

        [HttpGet("{id}")]
        public Book Get(int id)
        {
            var book = JsonConvert.DeserializeObject<IList<Book>>(System.IO.File.ReadAllText("data.json")).FirstOrDefault(b => b.Id == id);
            if (book != null)
            {
                var imageArray = System.IO.File.ReadAllBytes(book.Thumbnail);
                var base64ImageRepresentation = Convert.ToBase64String(imageArray);
                book.Thumbnail = base64ImageRepresentation;
            }

            return book;
        }

        [HttpGet("legacy")]
        public async Task<IActionResult> GetLegacyAsync()
        {
            var client = new HttpClient();
            client.BaseAddress = new Uri(Environment.GetEnvironmentVariable("LegacyServiceUri"));
            string[] traceHeaders = { "x-request-id", "x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid", "x-b3-sampled", "x-b3-flags", "b3" };
            foreach (var header in traceHeaders)
            {
                if (Request.Headers.ContainsKey(header))
                {
                    client.DefaultRequestHeaders.Add(header, Request.Headers[header].First());
                }
            }

            using (var response = await client.GetAsync("books"))
            {
                return Ok(await response.Content.ReadAsStringAsync());
            }
        }

        [HttpGet("feeling-lucky")]
        public IActionResult FeelingLucky()
        {
            var rng = new Random();
            var luckyDip = rng.Next(1, 100);
            return luckyDip >= 80 ? Ok("You are lucky") : StatusCode(500, "Better luck next time");
        }
    }
}