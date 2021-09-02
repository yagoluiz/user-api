using AutoBogus;
using User.API.Requests;
using User.API.Responses;

namespace User.Unit.Tests.Builders
{
    public static class UsersBuilder
    {
        public static UsersSearchRequest UsersSearchRequest =>
            new AutoFaker<UsersSearchRequest>()
                .Generate();

        public static UsersPaginationResponse UsersPaginationResponse =>
            new AutoFaker<UsersPaginationResponse>()
                .Generate();
    }
}
