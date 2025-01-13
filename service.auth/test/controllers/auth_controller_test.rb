require "test_helper"

class AuthControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = users(:one)
  end

  test "should be able to login" do
    get auth_login_url(email: @user.email), as: :json
    assert_response :success
  end

  test "should be able to register" do
    assert_difference("User.count") do
      post auth_register_url, params: {
        user: { 
          # generate unique email to avoid conflicts
          email: "test_#{SecureRandom.hex(4)}@example.com",
          first_name: @user.first_name,
          last_name: @user.last_name,
        },
      }, as: :json
    end

    assert_response :created
  end
end
