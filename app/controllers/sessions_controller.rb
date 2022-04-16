class SessionsController < ApplicationController
  before_action :set_user, only: %i[ login register ]

  def login
    @user = User.new
  end

  def register
    @user = User.new
  end

  def create
    @user = User.find_by(email: params[:user][:email])
    if @user&.authenticate(params[:shortcode])
      session[:user_id] = @user.id
      redirect_to accounts_path
    else
      @user = User.new
      flash[:error] = "Invalid email or shortcode."
      render :login
    end
  end

  def new
    @user = User.new(user_params)

    respond_to do |format|
      if @user.save
        session[:user_id] = @user.id
        format.html { redirect_to accounts_path, notice: "User was successfully created." }
        format.json { render :show, status: :created, location: @user }
      else
        format.html { render :register, status: :unprocessable_entity }
        format.json { render json: @user.errors, status: :unprocessable_entity }
      end
    end
  end

  def create
    @user = User.find_by(email: params[:user][:email])
    if @user && @user.authenticate(params[:user][:password])
      session[:user_id] = @user.id
      redirect_to root_path
    else
      @user = User.new
      flash[:error] = "Invalid email or password"
      render :new
    end
  end

  def destroy
    session.clear
    redirect_to root_path
  end

  private

  def set_user
    @user = User.new
  end

  def user_params
    params.require(:user).permit \
      :name,
      :email
  end
end
