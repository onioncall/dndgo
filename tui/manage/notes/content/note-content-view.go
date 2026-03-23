package content

func (m NoteContentModel) View() string {
	if m.IsEditing {
		return m.ContentTextArea.View()
	} else {
		return m.ContentViewPort.View()
	}
}
