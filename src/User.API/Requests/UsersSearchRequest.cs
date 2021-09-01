using System.ComponentModel.DataAnnotations;

namespace User.API.Requests
{
    public class UsersSearchRequest
    {
        [Required] public string Query { get; set; }
        [MinLength(0)] public int From { get; set; }
        [MaxLength(15)] public int Size { get; set; } = 15;
    }
}
