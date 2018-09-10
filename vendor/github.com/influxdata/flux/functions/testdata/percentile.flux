from(bucket: "test")
    |> range(start: 2018-05-22T19:50:26Z)
    |> group(by: ["_measurement", "_start"])
    |> percentile(percentile:0.75, method:"exact_selector")
    |> map(fn: (r) => {_time: r._time, percentile: r._value})
    |> yield(name:"0")
