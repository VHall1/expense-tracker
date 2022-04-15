class User < ApplicationRecord
  has_many :accounts, dependent: :destroy

  validates :name, presence: true
  validates :email, presence: true, uniqueness: true, format: { with: URI::MailTo::EMAIL_REGEXP }

  after_save :validate_email

  def balance
    self.accounts.sum(:balance)
  end

  def validate_email
    # Only try to validate email if we're changing
    # the user email address.
    return unless saved_change_to_email_validated?

    # TODO: Make this logic work
    # Maybe save token to Redis, since this should be shortlived
    # if email_validation_token.nil?
    #   self.email_validation_token = SecureRandom.uuid
    # end

    # UserMailer.validate_email(self).deliver_now
  end
end
