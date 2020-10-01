# pod-tzzh-mail

A [pod](https://github.com/babashka/babashka.pods) to send emails with [babashka](https://github.com/borkdude/babashka/)

It's relying entirely on https://github.com/jordan-wright/email and so supports most of its features

## Usage

``` clojure
(require '[babashka.pods])
(babashka.pods/load-pod ["./pod-tzzh-mail"])
(require '[pod.tzzh.mail :as m])

(s/send-mail {:host "smtp.gmail.com"
              :port 587
              :username "kylian.mbappe@gmail.com"
              :password "kylian123"
              :subject "Subject of the email"
              :from "kylian.mbappe@gmail.com"
              :to ["somebody@somehwere.com"]
              :cc ["somebodyelse@somehwere.com"]
              :text "aaa" ;; for text body
              :html "<b> kajfhajkfhakjs </b>" ;; for html body
              :attachments ["./do-everything.clj"] ;; paths to the files to attch
              })
```

To use with gmail you need to [allow less secure apps](https://myaccount.google.com/lesssecureapps)

## Debugging

For debugging set the environment variable `POD_TZZH_MAIL_DEBUG=true` and the logs will show in stderr.
