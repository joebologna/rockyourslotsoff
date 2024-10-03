@test "Spin" {
    pkill slots || true
    go generate ./...
    go run . &
    sleep 2
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Spin | jq '.Success') = "true" ]]
}

@test "UpdateBalance" {
    [[ $(curl -d '{"Amount": 100}' http://localhost:8998/VSlotService.UpdateBalance | jq '.Amount') -eq 100 ]]
}

@test "GetBalance" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.GetBalance | jq '.Amount') -eq 100 ]]
}

@test "Reset" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Reset | jq '.Success') = "true" ]]
}

@test "GetBalance2" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.GetBalance | jq '.Success') = "true" ]]
    pkill slots || true
}
