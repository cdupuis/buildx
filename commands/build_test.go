package commands

import (
	"os"
	"strings"
	"testing"
)

func setupTest(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
		os.Unsetenv("BUILDX_GIT_INFO")
	}
}

func TestAddGitProvenanceDataWithoutEnv(t *testing.T) {
	labels, err := addGitProvenance(nil, ".", "")
	if err != nil {
		t.Error("No error expected")
	}
	if labels != nil {
		t.Error("No labels expected")
	}
}

func TestAddGitProvenanceDataWithoutLabels(t *testing.T) {
	os.Setenv("BUILDX_GIT_INFO", "full")
	labels, err := addGitProvenance(nil, ".", "")
	if err != nil {
		t.Error("No error expected")
	}
	if len(labels) != 3 {
		t.Error("Exactly 3 git provenance labels expected")
	}
	dockerfileLabel := strings.Split(labels[2], "=")
	if dockerfileLabel[0] != "com.docker.image.dockerfile.path" || dockerfileLabel[1] != "Dockerfile" {
		t.Error("Expected a dockerfile path provenance label")
	}
	shaLabel := strings.Split(labels[0], "=")
	if shaLabel[0] != "org.opencontainers.image.revision" {
		t.Error("Expected a sha provenance label")
	}
	originLabel := strings.Split(labels[1], "=")
	if originLabel[0] != "org.opencontainers.image.source" {
		t.Error("Expected a origin provenance label")
	}
}

func TestAddGitProvenanceDataWithLabels(t *testing.T) {
	os.Setenv("BUILDX_GIT_INFO", "full")
	existingLabels := []string{"foo=bar"}
	labels, err := addGitProvenance(existingLabels, ".", "")
	if err != nil {
		t.Error("No error expected")
	}
	if len(labels) != 4 {
		t.Error("Exactly 3 git provenance labels expected")
	}
	dockerfileLabel := strings.Split(labels[3], "=")
	if dockerfileLabel[0] != "com.docker.image.dockerfile.path" || dockerfileLabel[1] != "Dockerfile" {
		t.Error("Expected a dockerfile path provenance label")
	}
	shaLabel := strings.Split(labels[1], "=")
	if shaLabel[0] != "org.opencontainers.image.revision" {
		t.Error("Expected a sha provenance label")
	}
	originLabel := strings.Split(labels[2], "=")
	if originLabel[0] != "org.opencontainers.image.source" {
		t.Error("Expected a origin provenance label")
	}
}
