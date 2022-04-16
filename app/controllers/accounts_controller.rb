class AccountsController < ApplicationController
  before_action :set_account, only: %i[ show edit update destroy ]

  # GET /accounts
  def index
    @accounts = current_user.accounts
  end

  # GET /accounts/1
  def show
  end

  # GET /accounts/new
  def new
    @account = Account.new
  end

  # GET /accounts/1/edit
  def edit
  end

  # POST /accounts
  def create
    # Only allow a list of trusted account types through.
    choices = Account::CHOICES.map { |_, account_type| account_type }
    unless choices.include? account_params[:type]
      flash[:error] = 'Invalid account type'
      redirect_to new_account_path
      return
    end

    @account = Account.new(account_params)
    @account.user = current_user
    # TODO: Make this an initial transaction?
    @account.balance = params[:initial_balance]&.to_f || 0.0

    respond_to do |format|
      if @account.save
        format.html { redirect_to accounts_url, notice: "Account was successfully created." }
        format.json { render :show, status: :created, location: @account }
      else
        format.html { render :new, status: :unprocessable_entity }
        format.json { render json: @account.errors, status: :unprocessable_entity }
      end
    end
  end

  # PATCH/PUT /accounts/1
  def update
    respond_to do |format|
      if @account.update(account_params)
        format.html { redirect_to account_url(@account), notice: "Account was successfully updated." }
        format.json { render :show, status: :ok, location: @account }
      else
        format.html { render :edit, status: :unprocessable_entity }
        format.json { render json: @account.errors, status: :unprocessable_entity }
      end
    end
  end

  # DELETE /accounts/1
  def destroy
    @account.destroy

    respond_to do |format|
      format.html { redirect_to accounts_url, notice: "Account was successfully destroyed." }
      format.json { head :no_content }
    end
  end

  private
    def set_account
      @account = Account.find(params[:id])
    end

    # Only allow a list of trusted parameters through.
    def account_params
      params.fetch(:account, {}).permit \
      :name,
      :type
    end
end
