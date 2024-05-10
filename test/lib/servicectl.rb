# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

module Servicectl
  class << self
    def client_config
      @client_config ||= Nonnative.configurations('.config/client.yml')
    end
  end
end
