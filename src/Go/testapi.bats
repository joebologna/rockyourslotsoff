@test "Spin" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Reset | jq '.Success') = "true" ]]
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Spin | jq '.Success') = "true" ]]
}

@test "TwoSpins" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Reset | jq '.Success') = "true" ]]
    spin=$(curl -d '{}' http://localhost:8998/VSlotService.Spin)
    [[ $(echo $spin | jq '.Success') = "true" ]]
    [[ $(echo $spin | jq '.Reels[]' | awk '{printf "%s", $0}') = "777" ]]
}

@test "UpdateBalance" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Reset | jq '.Success') = "true" ]]
    spin=$(curl -d '{}' http://localhost:8998/VSlotService.Spin)
    [[ $(echo $spin | jq '.Success') = "true" ]]
    [[ $(echo $spin | jq '.Reels[]' | awk '{printf "%s", $0}') = "777" ]]
    balance=$(curl -d '{"Amount":10}' http://localhost:8998/VSlotService.UpdateBalance | jq '.Amount')
    [[ $balance -eq 10 ]]
}

@test "TestLoser" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Reset | jq '.Success') = "true" ]]
    spin=$(curl -d '{}' http://localhost:8998/VSlotService.Spin)
    [[ $(echo $spin | jq '.Success') = "true" ]]
    [[ $(echo $spin | jq '.Reels[]' | awk '{printf "%s", $0}') = "777" ]]
    spin=$(curl -d '{}' http://localhost:8998/VSlotService.Spin)
    [[ $(echo $spin | jq '.Success') = "true" ]]
    [[ $(echo $spin | jq '.Reels[]' | awk '{printf "%s", $0}') = "452" ]]
    [[ $(echo $spin | jq '.IsWinner') = "false" ]]
}
