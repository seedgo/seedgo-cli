package action

import (
	"os"
	"testing"
)

// TestCreateProject_EmptyProjectName tests the case when project name is empty
func TestCreateProject_EmptyProjectName(t *testing.T) {
	projectName := ""
	err := createProject(projectName)

	// Verify that the function returns nil and prints the correct message
	if err.Error() != "please specifiy the new ProjectName" {
		t.Error("Expected error message to contain 'please specifiy the new ProjectName'")
	}
}

func TestCreateProject_ValidProjectName(t *testing.T) {
	// Create a temp dir as workdir for testing.
	workDir, err := os.MkdirTemp("", "test-")
	defer os.RemoveAll(workDir)
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	os.Chdir(workDir)
	println("workDir: %s", workDir)
	projectName := "test-project"
	err = createProject(projectName)

	// Verify that the function returns nil
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
