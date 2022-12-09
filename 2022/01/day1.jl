function copy!(array, start, elements)
    start_array = array[start:length(array)]
    j = 1
    for (i, _) in enumerate(start_array)
        array[i+start-1] = elements[j]
        j = j + 1
    end
end


function main(filename)
    lines = readlines(filename)

    top3 = [0, 0, 0]
    inventory = 0
    for (_, line) in enumerate(lines)
        if isempty(line)
            for (i, n) in enumerate(top3)
                if inventory > n
                    max = length(top3)
                    if i < max
                        # Shift lower inventories down
                        copy!(top3, i + 1, top3[i:max])
                    end
                    top3[i] = inventory
                    break
                end
            end
            inventory = 0
        else
            inventory += parse(Int, line)
        end
    end

    res1 = top3[1]
    res2 = sum(+, top3)

    println("part 1: ", res1)
    println("part 2: ", res2)
end

println("Test:")
main("test.txt")
println("Real:")
main("input.txt")
