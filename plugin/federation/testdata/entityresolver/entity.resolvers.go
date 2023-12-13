package entityresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/trevor-scheer/gqlgen version v0.17.41-dev

import (
	"context"
	"fmt"

	"github.com/trevor-scheer/gqlgen/plugin/federation/testdata/entityresolver/generated"
	"github.com/trevor-scheer/gqlgen/plugin/federation/testdata/entityresolver/generated/model"
)

// FindHelloByName is the resolver for the findHelloByName field.
func (r *entityResolver) FindHelloByName(ctx context.Context, name string) (*model.Hello, error) {
	return &model.Hello{
		Name: name,
	}, nil
}

// FindHelloMultiSingleKeysByKey1AndKey2 is the resolver for the findHelloMultiSingleKeysByKey1AndKey2 field.
func (r *entityResolver) FindHelloMultiSingleKeysByKey1AndKey2(ctx context.Context, key1 string, key2 string) (*model.HelloMultiSingleKeys, error) {
	panic(fmt.Errorf("not implemented"))
}

// FindHelloWithErrorsByName is the resolver for the findHelloWithErrorsByName field.
func (r *entityResolver) FindHelloWithErrorsByName(ctx context.Context, name string) (*model.HelloWithErrors, error) {
	if name == "inject error" {
		return nil, generated.ErrResolvingHelloWithErrorsByName
	} else if name == "" {
		return nil, generated.ErrEmptyKeyResolvingHelloWithErrorsByName
	}

	return &model.HelloWithErrors{
		Name: name,
	}, nil
}

// FindManyMultiHelloByNames is the resolver for the findManyMultiHelloByNames field.
func (r *entityResolver) FindManyMultiHelloByNames(ctx context.Context, reps []*model.MultiHelloByNamesInput) ([]*model.MultiHello, error) {
	results := []*model.MultiHello{}

	for _, item := range reps {
		results = append(results, &model.MultiHello{
			Name: item.Name + " - from multiget",
		})
	}

	return results, nil
}

// FindManyMultiHelloMultipleRequiresByNames is the resolver for the findManyMultiHelloMultipleRequiresByNames field.
func (r *entityResolver) FindManyMultiHelloMultipleRequiresByNames(ctx context.Context, reps []*model.MultiHelloMultipleRequiresByNamesInput) ([]*model.MultiHelloMultipleRequires, error) {
	results := make([]*model.MultiHelloMultipleRequires, len(reps))

	for i := range reps {
		results[i] = &model.MultiHelloMultipleRequires{
			Name: reps[i].Name,
		}
	}

	return results, nil
}

// FindManyMultiHelloRequiresByNames is the resolver for the findManyMultiHelloRequiresByNames field.
func (r *entityResolver) FindManyMultiHelloRequiresByNames(ctx context.Context, reps []*model.MultiHelloRequiresByNamesInput) ([]*model.MultiHelloRequires, error) {
	results := make([]*model.MultiHelloRequires, len(reps))

	for i := range reps {
		results[i] = &model.MultiHelloRequires{
			Name: reps[i].Name,
		}
	}

	return results, nil
}

// FindManyMultiHelloWithErrorByNames is the resolver for the findManyMultiHelloWithErrorByNames field.
func (r *entityResolver) FindManyMultiHelloWithErrorByNames(ctx context.Context, reps []*model.MultiHelloWithErrorByNamesInput) ([]*model.MultiHelloWithError, error) {
	return nil, fmt.Errorf("error resolving MultiHelloWorldWithError")
}

// FindManyMultiPlanetRequiresNestedByNames is the resolver for the findManyMultiPlanetRequiresNestedByNames field.
func (r *entityResolver) FindManyMultiPlanetRequiresNestedByNames(ctx context.Context, reps []*model.MultiPlanetRequiresNestedByNamesInput) ([]*model.MultiPlanetRequiresNested, error) {
	worlds := map[string]*model.World{
		"earth": {
			Foo: "A",
		},
		"mars": {
			Foo: "B",
		},
	}

	results := make([]*model.MultiPlanetRequiresNested, len(reps))

	for i := range reps {
		name := reps[i].Name
		world, ok := worlds[name]
		if !ok {
			return nil, fmt.Errorf("unknown planet: %s", name)
		}

		results[i] = &model.MultiPlanetRequiresNested{
			Name:  name,
			World: world,
		}
	}

	return results, nil
}

// FindPlanetMultipleRequiresByName is the resolver for the findPlanetMultipleRequiresByName field.
func (r *entityResolver) FindPlanetMultipleRequiresByName(ctx context.Context, name string) (*model.PlanetMultipleRequires, error) {
	return &model.PlanetMultipleRequires{Name: name}, nil
}

// FindPlanetRequiresByName is the resolver for the findPlanetRequiresByName field.
func (r *entityResolver) FindPlanetRequiresByName(ctx context.Context, name string) (*model.PlanetRequires, error) {
	return &model.PlanetRequires{
		Name: name,
	}, nil
}

// FindPlanetRequiresNestedByName is the resolver for the findPlanetRequiresNestedByName field.
func (r *entityResolver) FindPlanetRequiresNestedByName(ctx context.Context, name string) (*model.PlanetRequiresNested, error) {
	worlds := map[string]*model.World{
		"earth": {
			Foo: "A",
		},
		"mars": {
			Foo: "B",
		},
	}
	world, ok := worlds[name]
	if !ok {
		return nil, fmt.Errorf("unknown planet: %s", name)
	}

	return &model.PlanetRequiresNested{
		Name:  name,
		World: world,
	}, nil
}

// FindWorldByHelloNameAndFoo is the resolver for the findWorldByHelloNameAndFoo field.
func (r *entityResolver) FindWorldByHelloNameAndFoo(ctx context.Context, helloName string, foo string) (*model.World, error) {
	return &model.World{
		Hello: &model.Hello{
			Name: helloName,
		},
		Foo: foo,
	}, nil
}

// FindWorldNameByName is the resolver for the findWorldNameByName field.
func (r *entityResolver) FindWorldNameByName(ctx context.Context, name string) (*model.WorldName, error) {
	return &model.WorldName{
		Name: name,
	}, nil
}

// FindWorldWithMultipleKeysByHelloNameAndFoo is the resolver for the findWorldWithMultipleKeysByHelloNameAndFoo field.
func (r *entityResolver) FindWorldWithMultipleKeysByHelloNameAndFoo(ctx context.Context, helloName string, foo string) (*model.WorldWithMultipleKeys, error) {
	return &model.WorldWithMultipleKeys{
		Hello: &model.Hello{
			Name: helloName,
		},
		Foo: foo,
	}, nil
}

// FindWorldWithMultipleKeysByBar is the resolver for the findWorldWithMultipleKeysByBar field.
func (r *entityResolver) FindWorldWithMultipleKeysByBar(ctx context.Context, bar int) (*model.WorldWithMultipleKeys, error) {
	return &model.WorldWithMultipleKeys{
		Bar: bar,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
