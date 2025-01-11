package feed

import (
	"context"
	"io"
	"log"
	"os"
	"testing"
)

func setupTestLogging(t *testing.T) {
	t.Helper()
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	if testing.Verbose() {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}

func TestIntegrationWithJetBrainsFeedRelease(t *testing.T) {
	setupTestLogging(t)

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	feedURL := "https://download.jetbrains.com/toolbox/feeds/v1/release.feed.xz.signed"

	ctx := context.Background()
	err := downloadAndProcessFeed(ctx, feedURL)
	if err != nil {
		t.Fatalf("Failed to process JetBrains feed: %v", err)
	}

}

func TestIntegrationWithJetBrainsFeedEnterprise(t *testing.T) {
	setupTestLogging(t)

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	feedURL := "https://download.jetbrains.com/toolbox/feeds/v1/enterprise.feed.xz.signed"

	err := downloadAndProcessFeed(context.Background(), feedURL)
	if err != nil {
		t.Fatalf("Failed to process JetBrains feed: %v", err)
	}

}

func TestIntegrationWithJetBrainsFeedArm(t *testing.T) {
	setupTestLogging(t)

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	feedURL := "https://download.jetbrains.com/toolbox/feeds/v1/public-feed-arm.feed.xz.signed"

	err := downloadAndProcessFeed(context.Background(), feedURL)
	if err != nil {
		t.Fatalf("Failed to process JetBrains feed: %v", err)
	}

}

func TestIntegrationWithJetBrainsFeedAndroid(t *testing.T) {
	setupTestLogging(t)

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	feedURL := "https://download.jetbrains.com/toolbox/feeds/v1/android-studio.feed.xz.signed"

	err := downloadAndProcessFeed(context.Background(), feedURL)
	if err != nil {
		t.Fatalf("Failed to process JetBrains feed: %v", err)
	}

}
