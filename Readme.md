# Email Verification Guide


## 📌 Domain Knowledge

* An email follows the format: `username@domain.com`.
* The **domain** (e.g., `gmail.com`) must be properly structured and resolvable.
* The **username** must follow RFC 5322 syntax rules.
* Using DNS, computers/servers locate each other.
* **MX records** indicate which server accepts mail for a domain.

👉 If MX lookup fails, the email is marked as **unverified**.

---

## 📡 SMTP and Mailbox Verification

* Email servers connect via **SMTP (Simple Mail Transfer Protocol)**.
* Default connection is on **port 25**, but ISPs often block this for end-users.
* Alternative ports:

  * **465 (SMTPS)** and **587 (Submission Port)** → require authentication.

### Why Port 25 is risky

* Open to anyone, no authentication required.
* Spammers can flood the internet with spoofed emails.
* ISPs block it for normal users; only legit mail servers use it.

### Why Ports 465/587 are safer

* Require login (username/password, TLS).
* Providers (Google, Microsoft, Zoho, etc.) enforce:

  * Rate limits
  * Spam detection
  * Account suspension if abuse occurs

👉 Example:
If you configure Outlook with `hameed@gmail.com` over port **587**:

* Gmail requires authentication.
* Gmail checks for spam/abuse.
* Bulk spam attempts get blocked.

---

## 📬 Mailbox Existence Check

* After resolving MX records, SMTP commands can be used to check if a mailbox exists.
* Mail servers respond with status codes indicating success, failure, or restrictions.

---

## 🛑 Email Address Categories

### 1. Disposable Domains

* **Definition:** Temporary, single-use emails (e.g., `10minutemail.com`, `mailinator.com`).
* **Risk:** Often used for fake or short-term signups.
* **Action:** ❌ Reject.

---

### 2. Role-Based Accounts

* **Definition:** Functional addresses not tied to a single person (e.g., `info@`, `admin@`, `support@`).
* **Risk:** Shared access, not personal, unsuitable for unique user verification.
* **Action:** ⚠️ Flag or reject (based on use case).

---

### 3. Free Providers

* **Definition:** Public services like Gmail, Yahoo, Outlook, ProtonMail.
* **Risk:** Easy to create multiple accounts; potential for abuse.
* **Reality:** Very common and usually valid.
* **Action:** ✅ Allow, but consider encouraging business emails for B2B apps.

---

## ✅ Summary

* **Disposable → Reject**
* **Role-based → Flag**
* **Free provider → Allow (note: personal, not business)**

