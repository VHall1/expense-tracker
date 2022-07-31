class WebhooksController < ApplicationController
  def starling_feed_item
    CreateTransaction.process_async(params)
  end
end
