package models 
type notificationConfiguration struct {
	emailField []EmailEventConfig

	jiraField []JiraEventConfig

	webhookField []WebhookEventConfig
}

