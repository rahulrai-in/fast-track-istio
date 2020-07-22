using System;

namespace legacy
{
    public class Book
    {
        public int Id { get; set; }
        public DateTime Date { get; set; }
        public string Title { get; set; }
        public string Thumbnail { get; set; }
        public string Author { get; set; }
        public string Genre { get; set; }
        public string Summary { get; set; }
    }
}