class WebhooksController < ApplicationController
  def starling_feed_item(params)
    CreateTransaction.process_async(params.to_h)
  end
end
