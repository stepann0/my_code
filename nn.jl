using Random

sizes = [2, 3, 5, 9]
weights = [rand(sizes[i+1], sizes[i]) for i in 1:length(sizes)-1]
biases = [rand(i, 1) for i in sizes[2:end]]


relu(x) = max(0, x)
function feedforward(a)
    for (b, w) in zip(biases, weights)
        a = relu.(w*a+b)
    end
    return a
end

function SGD(training_data, epochs::Integer, 
             mini_batch_size::Integer, η::Float64, test_data=false)

    if test_data != false
        n_test = length(test_data)
    end

    n = length(training_data)
    for j in 1:epochs
        shuffle!(training_data)
        mini_batches = [training_data[k:k+15] 
                        for k in range(start=1, step=mini_batch_size, stop=n)
                        if k+15<=n]
        for mini_batch in mini_batches
            update_mini_batch(mini_batch, η)
        end
    end
end

function update_mini_batch(mini_batch, η)
    ∇b = [zeros(size(b)) for b in biases]
    ∇w = [zeros(size(w)) for w in weights]
    for (x, y) in mini_batch
        Δ∇b, Δ∇w = backprop(x, y)
        ∇b = [nb+dnb for (nb, dnb) in zip(∇b, Δ∇b)]
        ∇w = [nw+dnw for (nw, dnw) in zip(∇w, Δ∇w)]
    end

    self.weights = [w-(η/length(mini_batch))*nw 
                    for (w, nw) in zip(weights, ∇w)]
    self.biases = [b-(η/length(mini_batch))*nb 
                   for (b, nb) in zip(biases, ∇b)]
    end
td = [i for i in 1:200]
SGD(td, 1, 10, 0.1)