package config

import "testing"

func TestNormalizeConfigFormValuesFallsBackToValidDefaults(t *testing.T) {
	model, shell, suggest, explain := normalizeConfigFormValues(
		[]string{"llama3", "qwen2.5"},
		"missing-model",
		"fish",
		"",
		"experimental",
	)

	if model != "llama3" {
		t.Fatalf("expected first available model fallback, got %q", model)
	}
	if shell != defaultShell {
		t.Fatalf("expected default shell fallback, got %q", shell)
	}
	if suggest != defaultSuggestionPolicy {
		t.Fatalf("expected default suggestion policy fallback, got %q", suggest)
	}
	if explain != defaultExplainPolicy {
		t.Fatalf("expected default explain policy fallback, got %q", explain)
	}
}

func TestNormalizeConfigFormValuesKeepsValidSelections(t *testing.T) {
	model, shell, suggest, explain := normalizeConfigFormValues(
		[]string{"llama3", "qwen2.5"},
		"qwen2.5",
		"zsh",
		Creative,
		Stable,
	)

	if model != "qwen2.5" || shell != "zsh" || suggest != Creative || explain != Stable {
		t.Fatalf("expected values to remain unchanged, got model=%q shell=%q suggest=%q explain=%q", model, shell, suggest, explain)
	}
}
