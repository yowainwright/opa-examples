package concert

test_allow_entry_orange {
	input := {"wristband_color": "orange"}
	results := data.concert.privileges with input as input
	results.allow_entry == true
}

test_allow_vip_access_orange {
	input := {"wristband_color": "orange"}
	results := data.concert.privileges with input as input
	results.allow_vip_access == false
}

test_allow_free_food_orange {
	input := {"wristband_color": "orange"}
	results := data.concert.privileges with input as input
	results.allow_free_food == false
}

test_allow_entry_blue {
	input := {"wristband_color": "blue"}
	results := data.concert.privileges with input as input
	results.allow_entry == true
}

test_allow_vip_access_blue {
	input := {"wristband_color": "blue"}
	results := data.concert.privileges with input as input
	results.allow_vip_access == false
}

test_allow_free_food_blue {
	input := {"wristband_color": "blue"}
	results := data.concert.privileges with input as input
	results.allow_free_food == true
}

test_allow_entry_purple {
	input := {"wristband_color": "purple"}
	results := data.concert.privileges with input as input
	results.allow_entry == true
}

test_allow_vip_access_purple {
	input := {"wristband_color": "purple"}
	results := data.concert.privileges with input as input
	results.allow_vip_access == true
}

test_allow_free_food_purple {
	input := {"wristband_color": "purple"}
	results := data.concert.privileges with input as input
	results.allow_free_food == true
}
