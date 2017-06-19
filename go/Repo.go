package main

import "fmt"

var currentId int

var images Images

// Give us some seed data
func init() {
    RepoCreateImage(Image{Name: "Write presentation"})
    RepoCreateImage(Image{Name: "Host meetup"})
}

func RepoFindImage(id int) Image {
    for _, t := range images {
        if t.Id == id {
            return t
        }
    }
    // return empty Image if not found
    return Image{}
}

func RepoCreateImage(t Image) Image {
    currentId += 1
    t.Id = currentId
    images = append(images, t)
    return t
}

func RepoDestroyImage(id int) error {
    for i, t := range images {
        if t.Id == id {
            images = append(images[:i], images[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Image with id of %d to delete", id)
}
