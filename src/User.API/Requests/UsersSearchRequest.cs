using System.ComponentModel.DataAnnotations;

namespace User.API.Requests
{
    public class UsersSearchRequest
    {
        /// <summary>
        /// Term search
        /// </summary>
        [Required]
        public string Query { get; set; }

        /// <summary>
        /// Pagination from
        /// </summary>
        public int From { get; set; } = 0;

        /// <summary>
        /// Pagination size
        /// </summary>
        public int Size { get; set; } = 15;
    }
}
