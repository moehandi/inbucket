package client

import "testing"

func TestClientV1ListMailbox(t *testing.T) {
	var want, got string

	c, err := NewV1(baseURLStr)
	if err != nil {
		t.Fatal(err)
	}
	mth := &mockHTTPClient{}
	c.client = mth

	// Method under test
	c.ListMailbox("testbox")

	want = "GET"
	got = mth.req.Method
	if got != want {
		t.Errorf("req.Method == %q, want %q", got, want)
	}

	want = baseURLStr + "/api/v1/mailbox/testbox"
	got = mth.req.URL.String()
	if got != want {
		t.Errorf("req.URL == %q, want %q", got, want)
	}
}

func TestClientV1GetMessage(t *testing.T) {
	var want, got string

	c, err := NewV1(baseURLStr)
	if err != nil {
		t.Fatal(err)
	}
	mth := &mockHTTPClient{}
	c.client = mth

	// Method under test
	c.GetMessage("testbox", "20170107T224128-0000")

	want = "GET"
	got = mth.req.Method
	if got != want {
		t.Errorf("req.Method == %q, want %q", got, want)
	}

	want = baseURLStr + "/api/v1/mailbox/testbox/20170107T224128-0000"
	got = mth.req.URL.String()
	if got != want {
		t.Errorf("req.URL == %q, want %q", got, want)
	}
}

func TestClientV1GetMessageSource(t *testing.T) {
	var want, got string

	c, err := NewV1(baseURLStr)
	if err != nil {
		t.Fatal(err)
	}
	mth := &mockHTTPClient{
		statusCode: 200,
		body:       "message source",
	}
	c.client = mth

	// Method under test
	source, err := c.GetMessageSource("testbox", "20170107T224128-0000")
	if err != nil {
		t.Fatal(err)
	}

	want = "GET"
	got = mth.req.Method
	if got != want {
		t.Errorf("req.Method == %q, want %q", got, want)
	}

	want = baseURLStr + "/api/v1/mailbox/testbox/20170107T224128-0000/source"
	got = mth.req.URL.String()
	if got != want {
		t.Errorf("req.URL == %q, want %q", got, want)
	}

	want = "message source"
	got = source.String()
	if got != want {
		t.Errorf("Source == %q, want: %q", got, want)
	}
}

func TestClientV1DeleteMessage(t *testing.T) {
	var want, got string

	c, err := NewV1(baseURLStr)
	if err != nil {
		t.Fatal(err)
	}
	mth := &mockHTTPClient{}
	c.client = mth

	// Method under test
	c.DeleteMessage("testbox", "20170107T224128-0000")

	want = "DELETE"
	got = mth.req.Method
	if got != want {
		t.Errorf("req.Method == %q, want %q", got, want)
	}

	want = baseURLStr + "/api/v1/mailbox/testbox/20170107T224128-0000"
	got = mth.req.URL.String()
	if got != want {
		t.Errorf("req.URL == %q, want %q", got, want)
	}
}

func TestClientV1PurgeMailbox(t *testing.T) {
	var want, got string

	c, err := NewV1(baseURLStr)
	if err != nil {
		t.Fatal(err)
	}
	mth := &mockHTTPClient{}
	c.client = mth

	// Method under test
	c.PurgeMailbox("testbox")

	want = "DELETE"
	got = mth.req.Method
	if got != want {
		t.Errorf("req.Method == %q, want %q", got, want)
	}

	want = baseURLStr + "/api/v1/mailbox/testbox"
	got = mth.req.URL.String()
	if got != want {
		t.Errorf("req.URL == %q, want %q", got, want)
	}
}
