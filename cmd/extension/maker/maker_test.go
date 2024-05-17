package maker

import (
	"os"
	"testing"

	"github.com/gobeam/stringy"
)

func TestRenderTemplates(t *testing.T) {
	// setup temp dir output in ./testdata

	outputDir, err := os.MkdirTemp("", "TestEntity")
	if err != nil {
		t.Fatalf("Failed to create temp output directory: %v", err)
	}

	testEntityTemplatePath := "./testdata/Entity"
	outputEntityPath := outputDir + "/SwagFancyExtension/src/Content/TestEntity"
	namespace := "Swag\\FancyExtension"

	data := map[string]interface{}{
		"entityName":           "TestEntity",
		"tableName":            SnakeCase("TestEntity"),
		"parentClassNamespace": namespace,
	}

	err = renderTemplates(testEntityTemplatePath, outputEntityPath, data, true, namespace)
	if err != nil {
		t.Fatalf("Failed to render templates: %v", err)
	}

	// itterate through the output directory and check if the files are created
	// outputEntityPath has file TestEntityDefinition.php

	_, err = os.Stat(outputEntityPath + "/TestEntityDefinition.php")
	if err != nil {
		files, err := os.ReadDir(outputEntityPath)
		if err != nil {
			t.Fatalf("Failed to read directory %s: %v", outputEntityPath, err)
		}

		for _, file := range files {
			println(file.Name())
		}
		t.Fatalf("Failed to create TestEntityDefinition.php: %v", err)
	}

	defer os.RemoveAll(outputDir)
}

func TestPascalCaseHelper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "Test"},
		{"test entity", "TestEntity"},
		{"test-entity", "TestEntity"},
		{"test-entity_test", "TestEntityTest"},
		{"test_entity", "TestEntity"},
		{"TestEntity", "TestEntity"},
		{"testEntity", "TestEntity"},
		{"test_entity_definition", "TestEntityDefinition"},
	}

	for _, test := range tests {
		actual := PascalCase(test.input)
		if actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
		}
	}
}

func TestCamelCaseHelper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test entity", "testEntity"},
		{"test-entity", "testEntity"},
		{"test-entity_test", "testEntityTest"},
		{"test_entity", "testEntity"},
		{"TestEntity", "testEntity"},
		{"testEntity", "testEntity"},
		{"test_entity_definition", "testEntityDefinition"},
	}

	for _, test := range tests {
		actual := stringy.New(test.input).CamelCase().Get()
		if actual != test.expected {
			t.Errorf("Expected %s, got %s, input: %s", test.expected, actual, test.input)
		}
	}
}

func TestKebabCaseHelper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test entity", "test-entity"},
		{"test-entity", "test-entity"},
		{"test-entity_test", "test-entity-test"},
		{"test_entity", "test-entity"},
		{"testEntity", "test-entity"},
		{"TestEntity", "test-entity"},
		{"test_entity_definition", "test-entity-definition"},
	}

	for _, test := range tests {
		actual := KebabCase(test.input)
		if actual != test.expected {
			t.Errorf("Expected %s, got %s, input: %s", test.expected, actual, test.input)
		}
	}
}

func TestSnakeCaseHelper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test entity", "test_entity"},
		{"test-entity", "test_entity"},
		{"test-entity_test", "test_entity_test"},
		{"test_entity", "test_entity"},
		{"testEntity", "test_entity"},
		{"TestEntity", "test_entity"},
		{"test_entity_definition", "test_entity_definition"},
	}

	for _, test := range tests {
		actual := SnakeCase(test.input)
		if actual != test.expected {
			t.Errorf("Expected %s, got %s, input: %s", test.expected, actual, test.input)
		}
	}
}
