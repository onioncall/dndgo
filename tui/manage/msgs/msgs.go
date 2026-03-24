package msgs

type SetCurrentTabMsg struct{ Index int }

// Notes
type CharacterNotesUpdatedMsg struct{}
type NoteUpdatedMsg struct{}
type NoteSelectedMsg struct{}
type AddNoteMsg struct{ Title string }
type EditNoteMsg struct{}
type DeleteNoteMsg struct{}
