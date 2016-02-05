package server

import (
    "github.com/qmsk/e2/client"
    "fmt"
)

type Sources struct {
    client      *client.Client

    sourceMap   map[string]Source
}

func (sources *Sources) update() error {
    clientSources, err := sources.client.ListSources()
    if err != nil {
        return err
    }

    sources.sourceMap = make(map[string]Source)

    for _, source := range clientSources {
        source := Source{
            ID:     source.ID,
            Name:   source.Name,
        }

        sources.sourceMap[source.String()] = source
    }

    return nil
}

func (sources *Sources) Index(name string) (apiResource, error) {
    if err := sources.update(); err != nil  {
        panic(err)
    } else if source, found := sources.sourceMap[name]; !found {
        return nil, nil
    } else {
        return source, nil
    }
}

func (sources *Sources) Get() (interface{}, error) {
    if err := sources.update(); err != nil {
        return nil, err
    } else {
       return sources.sourceMap, nil
    }
}

type Source struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
}

func (source Source) String() string {
    return fmt.Sprintf("%d", source.ID)
}

func (source Source) Get() (interface{}, error) {
    return source, nil
}