package db

import (
	"context"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomWeightEntry(t *testing.T) Weightentry {
	user := CreateRandomUser(t)

	arg := CreateWeightEntryParams{
		Username:  user.Username,
		Notes:     util.GetRandomUsername(79),
		WeightKg:  int32(util.GetRandomAmount(1, 200)),
		EntryDate: time.Now(),
	}

	entry, err := testStore.CreateWeightEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Username, entry.Username)
	require.Equal(t, arg.Notes, entry.Notes)
	require.Equal(t, arg.WeightKg, entry.WeightKg)

	require.NotZero(t, entry.WeightEntryID)
	return entry
}

func TestCreateWeightEntry(t *testing.T) {
	CreateRandomWeightEntry(t)
}

func TestGetWeightEntry(t *testing.T) {
	entry1 := CreateRandomWeightEntry(t)
	arg := GetWeightEntryParams{
		WeightEntryID: entry1.WeightEntryID,
		Username:      entry1.Username,
	}
	entry2, err := testStore.GetWeightEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.Username, entry2.Username)
	require.Equal(t, entry1.Notes, entry2.Notes)
	require.Equal(t, entry1.WeightKg, entry2.WeightKg)
}

func TestUpdateWeightEntry(t *testing.T) {
	entry1 := CreateRandomWeightEntry(t)

	arg := UpdateWeightEntryParams{
		Username:      entry1.Username,
		WeightEntryID: entry1.WeightEntryID,
		Notes:         util.GetRandomUsername(79),
		WeightKg:      int32(util.GetRandomAmount(1, 200)),
		EntryDate:     time.Now(),
	}

	wkout2, err := testStore.UpdateWeightEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, arg.Notes, wkout2.Notes)
	require.Equal(t, arg.WeightKg, wkout2.WeightKg)
}

func TestDeleteWeightEntry(t *testing.T) {
	entry1 := CreateRandomWeightEntry(t)

	arg1 := DeleteWeightEntryParams{
		WeightEntryID: entry1.WeightEntryID,
		Username:      entry1.Username,
	}

	err := testStore.DeleteWeightEntry(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := GetWeightEntryParams{
		WeightEntryID: entry1.WeightEntryID,
		Username:      entry1.Username,
	}

	entry2, err := testStore.GetWeightEntry(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, entry2)
}

func TestListWeightEntries(t *testing.T) {
	lastEntry := CreateRandomWeightEntry(t)
	arg := ListWeightEntriesParams{
		Username: lastEntry.Username,
		Limit:    5,
		Offset:   0,
	}

	entries, err := testStore.ListWeightEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, lastEntry.Username, entry.Username)
	}
}
