# Simple linear regression in Julia & Flux.jl
# Простая линейная регрессия на Julia & Flux.jl
using Flux
using Flux: train!

actual(x) = 123.34x-231.98923
x_train, x_test = hcat(-1000:900...), hcat(900:1000...)
y_train, y_test = actual.(x_train), actual.(x_test)

model = Dense(1=>1)

loss(x, y) = Flux.Losses.mse(model(x), y)
opt = Descent()
data = [(x_train, y_train)]
parameters = Flux.params(model)

println("Training...")
for epoch in 1:1000
    train!(loss, parameters, data, opt)
    if epoch % 50 == 0
        println("Epoch $epoch... ℒ = $(loss(x_train, y_train))")
    end
end
