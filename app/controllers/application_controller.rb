class ApplicationController < ActionController::Base
  def current_user
    # @current_user ||= User.find(session[:user_id]) if session[:user_id]
    @current_user = User.first
  end
end
