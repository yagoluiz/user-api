using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using User.API.Requests;
using User.API.Responses;
using User.API.Services.Interfaces;

namespace User.API.Controllers
{
    [ApiController]
    [Produces("application/json")]
    [Route("search")]
    public class UsersController : ControllerBase
    {
        private readonly IUsersService _usersService;

        public UsersController(IUsersService usersService)
        {
            _usersService = usersService;
        }

        /// <summary>
        ///     Search users by term
        /// </summary>
        /// <remarks>
        ///     Sample request example:
        /// 
        ///     GET /search?query=yago (default: from = 0 and size = 15)
        /// 
        ///     GET /search?query=yago&amp;from=1&amp;size=10 (other from and size)
        /// </remarks>
        /// <param name="request"></param>
        /// <returns></returns>
        /// <response code="200">Users list</response>
        /// <response code="400">Bad request errors</response>
        /// <response code="500">Internal server error</response>
        [HttpGet]
        [ProducesResponseType(StatusCodes.Status200OK)]
        [ProducesResponseType(typeof(ProblemDetails), StatusCodes.Status400BadRequest)]
        [ProducesResponseType(typeof(ProblemDetails), StatusCodes.Status500InternalServerError)]
        public async Task<ActionResult<UsersPaginationResponse>> GetAllPaginationByTermAsync(
            [FromQuery] UsersSearchRequest request
        )
        {
            return Ok(await _usersService.GetAllPaginationByTermAsync(request));
        }
    }
}
