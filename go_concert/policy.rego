package concert

default privileges = {
	"allow_entry": false,
	"allow_vip_access": false,
	"allow_free_food": false,
}

privileges = {
	"allow_entry": true,
	"allow_vip_access": false,
	"allow_free_food": false,
} {
	input.wristband_color == "orange"
}

privileges = {
	"allow_entry": true,
	"allow_vip_access": false,
	"allow_free_food": true,
} {
	input.wristband_color == "blue"
}

privileges = {
	"allow_entry": true,
	"allow_vip_access": true,
	"allow_free_food": true,
} {
	input.wristband_color == "purple"
}
