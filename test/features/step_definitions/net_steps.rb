# frozen_string_literal: true

When('we try to setup http') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'http', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/http.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a successfully started http') do
  expect(@status.exitstatus).to eq(0)
end

When('we try to setup grpc') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../servicectl', 'grpc', '-i', 'file:.config/client.yml', '--verify')
  pid = spawn({}, cmd, %i[out err] => ['reports/grpc.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('we should have a successfully started grpc') do
  expect(@status.exitstatus).to eq(0)
end
