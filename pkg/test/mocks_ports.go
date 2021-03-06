package test

func (m *Mocks) ProvideSqsPort(name string) int {
	component := m.components[name].(*snsSqsComponent)
	return component.settings.SqsPort
}

func (m *Mocks) ProvideSnsPort(name string) int {
	component := m.components[name].(*snsSqsComponent)
	return component.settings.SnsPort
}

func (m *Mocks) ProvideCloudwatchPort(name string) int {
	component := m.components[name].(*cloudwatchComponent)
	return component.settings.Port
}

func (m *Mocks) ProvideDynamoDbPort(name string) int {
	component := m.components[name].(*dynamoDbComponent)
	return component.settings.Port
}

func (m *Mocks) ProvideElasticsearchPort(name string) int {
	component := m.components[name].(*elasticsearchComponent)
	return component.settings.Port
}

func (m *Mocks) ProvideKinesisPort(name string) int {
	component := m.components[name].(*kinesisComponent)
	return component.settings.Port
}

func (m *Mocks) ProvideMysqlPort(name string) int {
	component := m.components[name].(*mysqlComponent)
	return component.settings.Port
}

func (m *Mocks) ProvideRedisPort(name string) int {
	component := m.components[name].(*redisComponent)
	return component.settings.Port
}

func (m *Mocks) ProvideWiremockPort(name string) int {
	component := m.components[name].(*wiremockComponent)
	return component.settings.Port
}
