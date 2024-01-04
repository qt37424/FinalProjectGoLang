import { Component, OnInit, ViewChild } from "@angular/core";
import { NgForm } from "@angular/forms";

import { passwordService } from "../../services/password.service";
import { environment } from 'src/environments/environment';
import { Observable } from "rxjs";
import { map } from "rxjs/operators";

@Component({
  selector: "app-register",
  templateUrl: "./register.component.html",
  styleUrls: ["./register.component.css"],
  providers: [passwordService],
})
export class RegisterAccountComponent implements OnInit {
  @ViewChild("form", { static: false }) form: NgForm;

  submitted: boolean = false;
  email: string;
  message: string = "";
  isLoading = false;

  constructor(private passwordService: passwordService) {}

  ngOnInit() {
    // this.submitted = false;
  }

  /// Try to rewrite here to connect database from Golang -- QuangTT12: TODO
  onSubmit(resetForm: NgForm) {
    console.log("TRAN TRINH QUANG")
    var serchfind: boolean;

    this.email = resetForm.value.recoveremail;
    //EMAIL REGEX
    let regexp = new RegExp(
      /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
    //test user input
    serchfind = regexp.test(this.email);

    if (serchfind === true) {
      this.isLoading = true;
      this.sendRequestCreateAccount(this.email);
    } else {
      this.message = "Make sure you are entering correct email address";
      this.submitted = true;
    }
    // resetForm.reset();
  }

  private recover(email: string): void {
    this.passwordService.sendRecovery(email).subscribe(
      (response) => {
        console.log(response);
        if (response === "Request processed.") {
          this.message =
            "Your request has been processed. If the email matches our records, you will receive an email. Please check the Spam folder if the email response was not received within 12 hours.";
          this.form.reset();
          this.isLoading = false;
        } else {
          this.message =
            "We could not process your request. Make sure you are entering correct email address";
          this.isLoading = false;
        }
      },
      (error: any) => {
        this.message =
          "Unfortunately, we encountered an error. Please check your internet connection or try again later";
        this.form.reset();
        this.isLoading = false;
      }
    );
    this.submitted = true;
  }

  /// I need to rewrite it to createAccount
  private sendRequestCreateAccount(email: string): void {
    this.passwordService.sendRecovery(email).subscribe(
      (response) => {
        console.log(response);
        if (response === "Request processed.") {
          this.message =
            "Your request has been processed. If the email matches our records, you will receive an email. Please check the Spam folder if the email response was not received within 12 hours.";
          this.form.reset();
          this.isLoading = false;
        } else {
          this.message =
            "We could not process your request. Make sure you are entering correct email address";
          this.isLoading = false;
        }
      },
      (error: any) => {
        this.message =
          "Unfortunately, we encountered an error. Please check your internet connection or try again later";
        this.form.reset();
        this.isLoading = false;
      }
    );
    this.submitted = true;
  }
}
