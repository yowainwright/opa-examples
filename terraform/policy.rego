package main

import input as tfplan

resource_tags = {resource: tags |
	resource := tfplan.resource_changes[_]
	tags := resource.change.after.tags
	not resource.change.before.tags
}

deny[msg] {
	resource_tags[resource_name] = tags
	resource_type := split(resource_name, ".")[0]

	tags.example_tag == null
	msg := sprintf("Resource '%s' of type '%s' is missing the required 'example_tag'", [resource_name, resource_type])
}
