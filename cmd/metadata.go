package cmd

import (
	"fmt"
	fs "path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chazzychouse/venom/internal/ui"
	"github.com/spf13/cobra"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata [file]",
	Short: "Read and write audio file metadata",
	Long: `Read, write, and batch edit metadata (ID3 tags, etc.) for audio files.

Examples:
  venom metadata song.mp3              # Read metadata
  venom metadata song.mp3 --set-title "New Title"
  venom metadata ./music --batch       # Batch process directory`,
	Args: cobra.MinimumNArgs(1),
	RunE: runMetadata,
}

func init() {
	// TODO: Add flags for metadata operations
	// metadataCmd.Flags().StringP("set-title", "t", "", "set track title")
	// metadataCmd.Flags().StringP("set-artist", "a", "", "set artist name")
	// metadataCmd.Flags().BoolP("batch", "b", false, "batch process directory")
}
func doShit(p *tea.Program) {
	time.Sleep(2 * time.Second)
	p.Send(ui.SpinnerDone{})
}

func runMetadata(cmd *cobra.Command, args []string) error {
	filePath := args[0]

	s := ui.NewSpinner("Looking at your " + filePath + "...")
	p := tea.NewProgram(s)

	go func() {
		doShit(p)
	}()

	p.Run()
	var ext = fs.Ext(filePath)
	switch ext {
	case "mp3":
		fmt.Println("mp3")
	default:
	}
	// TODO: Implement metadata reading/writing
	// 1. Detect file type (mp3, flac, wav, etc.)
	// 2. Use appropriate library to read tags
	// 3. Display or modify based on flags
	//
	// Suggested libraries:
	//   - github.com/dhowden/tag (read-only, simple)
	//   - github.com/bogem/id3v2 (read/write ID3v2)

	return nil
}
