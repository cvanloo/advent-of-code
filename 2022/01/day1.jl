function copy!(array, start, elements)
    max = length(array)
    for (i, v) in enumerate(elements)
        idx = i+start-1
        if idx > max
            break
        end
        array[idx] = v
    end
end

# Both work, which one makes the intent more clear?
# Why isn't there anything like rusts enumerator.skip(n)?
function copy2!(array, start, elements)
    j = 1
    for (i, v) in enumerate(array)
        if i < start
            continue
        end
        array[i] = elements[j]
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
