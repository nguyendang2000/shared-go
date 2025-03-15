package redis

import (
	"fmt"
)

// SAdd adds one or more members to a set stored at the given key.
// If the key does not exist, a new set is created before adding the members.
func (inst *Service) SAdd(key string, members ...any) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	if err := inst.client.SAdd(ctx, key, members...).Err(); err != nil {
		return fmt.Errorf(ErrSAdd, err)
	}

	return nil
}

// SCard returns the number of members in the set stored at the given key.
// If the key does not exist, it is treated as an empty set and returns 0.
func (inst *Service) SCard(key string) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	card, err := inst.client.SCard(ctx, key).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrSCard, err)
	}

	return card, nil
}

// SDiff returns the difference between the set stored at key and the sets stored at the specified keys.
func (inst *Service) SDiff(key string, keys ...string) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	diff, err := inst.client.SDiff(ctx, append([]string{key}, keys...)...).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSDiff, err)
	}

	return diff, nil
}

// SDiffStore stores the difference between the set stored at key and the sets stored at the specified keys in the destination set.
func (inst *Service) SDiffStore(destination, key string, keys ...string) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	err := inst.client.SDiffStore(ctx, destination, append([]string{key}, keys...)...).Err()
	if err != nil {
		return fmt.Errorf(ErrSDiffStore, err)
	}

	return nil
}

// SInter returns the intersection of all the given sets.
func (inst *Service) SInter(key string, keys ...string) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	inter, err := inst.client.SInter(ctx, append([]string{key}, keys...)...).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSInter, err)
	}

	return inter, nil
}

// SInterCard returns the cardinality (size) of the intersection of all the given sets.
func (inst *Service) SInterCard(limit int64, key string, keys ...string) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	interCard, err := inst.client.SInterCard(ctx, limit, append([]string{key}, keys...)...).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrSInterCard, err)
	}

	return interCard, nil
}

// SInterStore stores the intersection of all the given sets in the destination set.
func (inst *Service) SInterStore(destination, key string, keys ...string) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	err := inst.client.SInterStore(ctx, destination, append([]string{key}, keys...)...).Err()
	if err != nil {
		return fmt.Errorf(ErrSInterStore, err)
	}

	return nil
}

// SIsMember checks if a given member exists in the set stored at the given key.
func (inst *Service) SIsMember(key string, member any) (bool, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	isMember, err := inst.client.SIsMember(ctx, key, member).Result()
	if err != nil {
		return false, fmt.Errorf(ErrSIsMember, err)
	}

	return isMember, nil
}

// SMembers returns all the members of the set stored at the given key.
func (inst *Service) SMembers(key string) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	members, err := inst.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSMembers, err)
	}

	return members, nil
}

// SMIsMember checks if multiple members exist in the set stored at the given key.
func (inst *Service) SMIsMember(key string, members ...any) ([]bool, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	isMember, err := inst.client.SMIsMember(ctx, key, members...).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSMIsMember, err)
	}

	return isMember, nil
}

// SMove moves a member from one set to another.
func (inst *Service) SMove(source, destination string, member any) (bool, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	moved, err := inst.client.SMove(ctx, source, destination, member).Result()
	if err != nil {
		return false, fmt.Errorf(ErrSMove, err)
	}

	return moved, nil
}

// SPop removes and returns one or more random members from the set stored at the given key.
func (inst *Service) SPop(key string, count int64) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	popped, err := inst.client.SPopN(ctx, key, count).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSPop, err)
	}

	return popped, nil
}

// SRandMember returns one or more random members from the set stored at the given key without removing them.
func (inst *Service) SRandMember(key string, count int64) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	members, err := inst.client.SRandMemberN(ctx, key, count).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSRandMember, err)
	}

	return members, nil
}

// SRem removes one or more members from the set stored at the given key.
func (inst *Service) SRem(key string, members ...any) (int64, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	removed, err := inst.client.SRem(ctx, key, members...).Result()
	if err != nil {
		return -1, fmt.Errorf(ErrSRem, err)
	}

	return removed, nil
}

// SUnion returns the union of all the given sets.
func (inst *Service) SUnion(key string, keys ...string) ([]string, error) {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	union, err := inst.client.SUnion(ctx, append([]string{key}, keys...)...).Result()
	if err != nil {
		return nil, fmt.Errorf(ErrSUnion, err)
	}

	return union, nil
}

// SUnionStore stores the union of all the given sets in the destination set.
func (inst *Service) SUnionStore(destination, key string, keys ...string) error {
	ctx, cancel := inst.getTimeout()
	defer cancel()

	err := inst.client.SUnionStore(ctx, destination, append([]string{key}, keys...)...).Err()
	if err != nil {
		return fmt.Errorf(ErrSUnionStore, err)
	}

	return nil
}
