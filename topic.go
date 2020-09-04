package gopic

type Any interface {}

type Listener func(any Any)

type TopicManager struct {
  listeners map[string][]Listener
}

func NewTopicManager() *TopicManager  {
  return &TopicManager{
    make(map[string][]Listener),
  }
}

func (m *TopicManager) Subscribe(topic string, listener Listener)  {
  if _, ok := m.listeners[topic]; !ok {
    m.listeners[topic] = []Listener {listener}
  } else {
    m.listeners[topic] = append(m.listeners[topic], listener)
  }
}

func (m *TopicManager) Publish(topic string, data Any)  {
  if listeners, ok := m.listeners[topic]; ok {
    for _, listener := range listeners  {
      listener(data)
    }
  }
}

var defaultManager *TopicManager

func init() {
  defaultManager = NewTopicManager()
}

func Publish(topic string, data Any)  {
 defaultManager.Publish(topic, data)
}

func Subscribe(topic string, listener Listener)  {
 defaultManager.Subscribe(topic, listener)
}
